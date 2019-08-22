echo compiling webpack
npm run compile
echo compiling binary
npm run go-compile
git add .
git commit -m "Deploy 8/21/2019"
git push origin deploy

ssh -T -i $key ubuntu@$stock_app_ec2 << 'ENDSSH'

exit

ENDSSH
