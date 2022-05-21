alias g := golden

default:
    @just --list

golden:
    tar -cz -f ./testdata/golden.tar.gz -C ./testdata golden
