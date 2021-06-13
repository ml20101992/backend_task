#go to app folder and run npm install
cd app
echo "Running npm install"
npm install
echo "npm install done"


#go to golang folder and run go install
cd ../processing_service
echo "Running go install"
go build
echo "go install done"

#create out folder for saving the files
cd ../
echo "Creating output folder"
rm -r out
mkdir out

echo "Setup done"
