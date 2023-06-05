
echo -e "bad example 00:\n"
go run . test/badexample00.txt
echo -e "bad example 01:\n"
go run . test/badexample01.txt
echo -e "bad example 02:\n"
go run . test/badexample02.txt
echo -e "bad example 03:\n"
go run . test/badexample03.txt
echo -e "bad example 04:\n"
go run . test/badexample04.txt
echo -e "bad format:\n"
go run . test/badformat.txt
echo "Press Enter to continue"
read -p "$*"
clear
echo -e "good example 00:\n"
go run . test/goodexample00.txt
echo "Press Enter to continue"
read -p "$*"
clear
echo -e "good example 01:\n"
go run . test/goodexample01.txt
echo "Press Enter to continue"
read -p "$*"
clear
echo -e "good example 02:\n"
go run . test/goodexample02.txt
echo "Press Enter to continue"
read -p "$*"
clear
echo -e "good example 03:\n"
go run . test/goodexample03.txt
echo "Press Enter to continue"
read -p "$*"
clear
echo -e "hard example:\n"
go run . test/hardexample.txt