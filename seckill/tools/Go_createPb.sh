# 进入pb文件目录
cd ../pb/

#!bin/sh
#if ! [ -e .././CsForClient/ ]
#then
#	mkdir .././CsForClient/
#fi

# 编译所有的proto文件
echo "Building proto files..."

for file in ./*
do
	# 检查是否是文件
	if [ -f $file ]
		then
			if [[ $file == *.proto ]]
				then
					# 如果为proto文件,则直接编译
					protoc -I=. --go_out=plugins=grpc:../pb/ ./*.proto 
			fi
	# 如果为文件夹
	elif [ -d $file ]
		then
			# 排除编译文件夹
			if [[ $file == ./WorldMap* ]]
			then
			    echo "Go proto Exclude Building  $file"
			    continue
			fi

			if ! [ -e ../pb/$file ]
				then
					# 不存在对应文件夹时新建
					mkdir ../pb/$file
			fi
			
			# 编译文件夹中所有的proto文件到指定文件夹
			protoc -I=. --go_out=plugins=grpc:../pb/  $file/*.proto
			
	fi
done
echo "Building proto files success"

# 如果有错误信息,则不退出界面

read -p "Press any key to continue"
