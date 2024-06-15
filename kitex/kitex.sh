kitex -module github.com/ulyssesorz/douyin -I ./ -v -service usersrv user.proto

kitex -module github.com/ulyssesorz/douyin -I ./ -v -service commentsrv comment.proto

kitex -module github.com/ulyssesorz/douyin -I ./ -v -service relationsrv relation.proto

kitex -module github.com/ulyssesorz/douyin -I ./ -v -service favoritesrv favorite.proto

kitex -module github.com/ulyssesorz/douyin -I ./ -v -service messagesrv message.proto

kitex -module github.com/ulyssesorz/douyin -I ./ -v -service videosrv video.proto
