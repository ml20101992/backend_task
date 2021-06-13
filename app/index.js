const FileLoader = require("./io/FileLoader.js");
const NatsService = require("./nats/NatsService.js");

//check if the file path was passed to the app
if (process.argv.length <= 2) {
  console.error("File parameter path not found.");
  return;
}

//get the file path
let filePath = process.argv[2];

//load the file
let file = FileLoader.loadFile(filePath);

//Check if file properly loaded
if (!file.success) {
  console.error("File not found.");
}

let nats = new NatsService();
nats
  .connect()
  .then(() => {
    nats.sendFile(file.data, file.name);
  })
  .then(() => {
      nats.consumeMessage();
  });
