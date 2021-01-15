docker rm -f $(docker ps -a |  grep "coreseek_main"  | awk '{print $1}')

docker run -it -d  \
	--restart=always \
	--name coreseek_main \
	-v $(pwd):/usr/local/etc/sphinx \
	-v $(pwd)/entry:/entry \
	-p 9312:9312 pastyouth/coreseek

