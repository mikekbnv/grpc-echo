package echo

func Echo(request *Msg) (*Reply, error) {
	return &Reply{Message: request.Text}, nil
}
