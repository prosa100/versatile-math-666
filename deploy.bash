git add --all :/
echo "Message: "
read msg
git commit --message="$msg"
git push origin master