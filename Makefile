pdf:
	pandoc --pdf-engine=xelatex -o hasil/buku-go--bpdp.pdf --toc --from markdown --template eisvogel --listings --top-level-division=chapter isi/*
