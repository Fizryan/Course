const {default: axios} = require("axios")

const text = "TOBAT!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!";
const url1 = `https://api.telegram.org/bot6645936098:AAEsKwAaoDnfsOOMwsH9dgbSjVcRV_nBar4/sendMessage?parse_mode=markdown&chat_id=6887431017&text=${text}`;

async function sendMessage() {
    try {
        while (true) {
            await axios.get(url1)
                .then((response) => {
                    if (response.status == 200) {
                        console.log("Oke")
                    }
                })
        }
    } catch (error) {
        console.error("API down!", error.message)
        sendMessage();
    }
}

sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();
sendMessage();