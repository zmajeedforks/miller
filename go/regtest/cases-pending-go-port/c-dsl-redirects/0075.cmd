mlr head -n 4 then put -q '@a[NR]=$a; @b[NR]=$b; emitp | "tr \[a-z\] \[A-Z\]", @*, "NR"' regtest/input/abixy
