package internal

func _doc(g *Generator) {
	//
	g.generateMessageStruct(nil, nil)

	// 提取注释
	extractComments(nil)

	newDescriptor(nil, nil, nil, 0)
	newEnumDescriptor(nil, nil, nil, 0)
}
