package common


// InitializeKubeClient by kubeconfig
// kubernetes client version is old, need upgrade
//func InitializeKubeClient(kubeconfigPath string) (*kubernetes.Clientset, error) {
//	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
//	if err != nil {
//		return nil, err
//	}
//	return kubernetes.NewForConfig(config)
//}