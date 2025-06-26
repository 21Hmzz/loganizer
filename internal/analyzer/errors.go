package analyzer

import "fmt"

type FileNotFoundErr struct{ Path string }

func (e *FileNotFoundErr) Error() string {
	return fmt.Sprintf("fichier ibntrouvable : %s", e.Path)
}

type ParseErr struct{ ID string }

func (e *ParseErr) Error() string {
	return fmt.Sprintf("erreur de pars sur : %s", e.ID)
}
