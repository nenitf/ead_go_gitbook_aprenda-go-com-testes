# ead_go_gitbook_aprenda-go-com-testes

Leitura do livro [Aprenda go com testes](https://larien.gitbook.io/aprenda-go-com-testes/)

```sh
go test ./folder

# todos testes
#go test ./...

# go test -bench="."
# go test -bench="./folder"

# go test -cover ./folder
```

## Notes

- É preferível que um usuário possa entender o uso de seu código apenas observando a assinatura de tipo e a documentação

### Links

- [Layout de um projeto](https://github.com/golang-standards/project-layout)
- [Post sobre estrutura de projeto](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
- [Naming](https://talks.golang.org/2014/names.slide#1)
    > Keep them short; long names obscure what the code does
- Documentação com `godoc -http :8000`
    - Instalar com `go get -v  golang.org/x/tools/cmd/godoc`
