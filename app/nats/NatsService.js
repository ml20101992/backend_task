const { connect, JSONCodec } = require("nats");
const jc = JSONCodec();

module.exports = class NatsService {
    constructor() {
        this.channel = "fileChannel"
    }

    async connect() {
        this.connection = await connect({ url: "nats://localhost:4222", json: true })
    }

    sendFile(buffer, filename) {
        let payload = {
            FileName: filename,
            FileData: buffer.toString("base64")
        }      

        console.log("Sending File.")
        this.connection.publish(this.channel, jc.encode(payload))
    }

    consumeMessage() {
        let sub = this.connection.subscribe(this.channel);
        console.log("Waiting for response");
        (async () => {
          for await (const m of sub) {
              let response = jc.decode(m.data);
              
              if (response.Success) {
                  console.log("Operation successful.");
                  console.log("File path: " + response.Path);
              } else {
                  console.error("Operation failed.\r\nError: " + response.Error)
              }

              await this.close();
          }
        })();
    }

    async close() {
        console.log("Closing connection");
        await this.connection.close();
    }
}