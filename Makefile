# 生成 swagger 文档
gen_docs:
	@echo "Build swagger docs"
	swag init -g ./router/router.go
