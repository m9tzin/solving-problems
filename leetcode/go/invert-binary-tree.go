package main

import "fmt"

// TreeNode representa um nó de uma árvore binária
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree inverte uma árvore binária recursivamente
// A ideia é trocar os filhos esquerdo e direito de cada nó
func invertTree(root *TreeNode) *TreeNode {
	// Caso base: se o nó é nil, retorna nil
	if root == nil {
		return nil
	}

	// Inverte recursivamente as subárvores esquerda e direita
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	// Troca os filhos esquerdo e direito
	root.Left = right
	root.Right = left

	return root
}

// printTreeInOrder imprime a árvore em ordem (para visualização)
func printTreeInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printTreeInOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	printTreeInOrder(root.Right)
}

// printTreePreOrder imprime a árvore em pré-ordem (para visualização)
func printTreePreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTreePreOrder(root.Left)
	printTreePreOrder(root.Right)
}

// createSampleTree cria uma árvore de exemplo para teste
func createSampleTree() *TreeNode {
	// Cria a árvore:
	//      4
	//    /   \
	//   2     7
	//  / \   / \
	// 1   3 6   9
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}
	return root
}

func main() {
	// Cria uma árvore de exemplo
	fmt.Println("=== Inversão de Árvore Binária ===")

	// Árvore original
	root := createSampleTree()
	fmt.Println("Árvore original (pré-ordem):")
	printTreePreOrder(root)
	fmt.Println()

	// Inverte a árvore
	invertedRoot := invertTree(root)
	fmt.Println("Árvore invertida (pré-ordem):")
	printTreePreOrder(invertedRoot)
	fmt.Println()

	// Verifica o resultado esperado
	fmt.Println("\nResultado esperado após inversão:")
	fmt.Println("4 7 9 6 2 3 1")
}
