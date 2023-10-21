package commands

import "fmt"

func Alias(enable bool, wbso_path string) error {
	aliasCommand := fmt.Sprintf("alias wbso='%sWBSO'", wbso_path)

	fmt.Println("To create the 'wbso' alias, add the following line to your .alias, .bash_aliases or .bash_profile:")
	fmt.Println(aliasCommand)
	return nil
}
