package cmds

import (
	"context"
	"errors"
	"fmt"

	"k8s.io/client-go/rest"

	"github.com/adhara-tech/kubesu/pkg/rpc"
	"github.com/adhara-tech/kubesu/pkg/types"

	"github.com/spf13/cobra"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var rpcPort string
var namespace string
var selectors string

var CmdConnectNodes = &cobra.Command{
	Use:   "connectNodes",
	Short: "Connect ",
	Args:  cobra.NoArgs,
	Run:   RunConnectNodes,
}

func InitCmdConnectNodes() {
	CmdConnectNodes.Flags().StringVarP(&namespace, "namespace", "", "", "Namespace")
	CmdConnectNodes.Flags().StringVarP(&selectors, "selector", "", "", "Label selector")
	CmdConnectNodes.Flags().StringVarP(&rpcPort, "rpcport", "", "", "RPC Port")
	CmdConnectNodes.MarkFlagRequired("namespace")
	CmdConnectNodes.MarkFlagRequired("selector")
	CmdConnectNodes.MarkFlagRequired("rpcport")
}

func RunConnectNodes(cmd *cobra.Command, args []string) {
	allNodes, err := retrieveNodes()
	if err != nil {
		panic(err)
	}

	fmt.Println()

	err = connectNodes(allNodes)
	if err != nil {
		panic(err)
	}
}

func retrieveNodes() ([]types.Node, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	k8s, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	k8sClient := k8s.CoreV1()
	pods, err := k8sClient.Pods(namespace).List(context.Background(), metav1.ListOptions{LabelSelector: selectors})
	if err != nil {
		return nil, err
	}

	var allNodes []types.Node
	for i, pod := range pods.Items {
		nodeIP := pod.Status.PodIP

		rpcClient, err := rpc.NewAdminApi(nodeIP, rpcPort)
		if err != nil {
			fmt.Printf("WARNING: %v\n", err)
			continue
		}

		nodeInfo, err := rpcClient.NodeInfo()
		if err != nil {
			fmt.Printf("WARNING: %v\n", err)
			continue
		}

		newNode := types.Node{
			Name:   pod.Name,
			Ip:     nodeIP,
			Port:   nodeInfo.Port,
			Pubkey: nodeInfo.Pubkey,
		}
		allNodes = append(allNodes, newNode)

		fmt.Printf("%d: %s\t%s\t%s\n", i, newNode.Ip, newNode.Name, newNode.Pubkey)
	}

	return allNodes, nil
}

// Loop through all nodes and add missing peers
func connectNodes(allNodes []types.Node) error {
	for _, node := range allNodes {
		fmt.Printf("Checking node [%s]\n", node.Name)
		err := connectMissingPeers(node, allNodes)
		if err != nil {
			fmt.Printf("WARNING: %v\n", err)
			continue
		}

		fmt.Println()
	}

	return nil
}

// Connect any missing peers of a single node
func connectMissingPeers(currentNode types.Node, allNodes []types.Node) error {
	rpcClient, err := rpc.NewAdminApi(currentNode.Ip, rpcPort)
	if err != nil {
		return err
	}

	peers, err := rpcClient.Peers()
	if err != nil {
		return err
	}

	for _, peerToAdd := range allNodes {
		// Exclude current node
		if currentNode.Ip == peerToAdd.Ip {
			continue
		}

		addPeerEnode := true
		for _, peer := range peers {
			if peerToAdd.Pubkey == peer.Pubkey {
				addPeerEnode = false
			}
		}

		if addPeerEnode {
			fmt.Printf("add:  %s\n", peerToAdd.Name)
			success, err := rpcClient.AddPeer(peerToAdd.Enode())
			if err != nil {
				return err
			}

			if !success {
				return errors.New(fmt.Sprintf("Failed to add [%s] to [%s]", peerToAdd.Name, currentNode.Name))
			}

		} else {
			fmt.Printf("has:  %s\n", peerToAdd.Name)
		}
	}

	return nil
}
