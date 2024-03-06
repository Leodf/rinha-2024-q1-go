tag=$(date +"%Y%m%d%H%M")

docker build -t rinha-2024q1-crebito-go:$tag -t leodf41/rinha-2024q1-crebito-go:$tag .
docker push leodf41/rinha-2024q1-crebito-go:$tag

echo $tag