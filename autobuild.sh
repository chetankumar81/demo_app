go build main.go
if [ $? -eq 0 ]; then
	rm main.zip
	zip main.zip main

	if [[ $__is_codebuild -ne 1 ]];
	then
		read -t 2 -e -p  "Abort local invocation? [y/n] " yn 
		if [[ $yn == "y" || $yn == "Y" ]]; 
		then	
			exit 0
		else
			echo ""
			sam local start-api
		fi
	fi
fi