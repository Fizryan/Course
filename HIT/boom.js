const axios = require("axios");

const token = "6645936098:AAEsKwAaoDnfsOOMwsH9dgbSjVcRV_nBar4";
const chatId = "6887431017";
const text = "TOBAT!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!";
const url = `https://api.telegram.org/bot${token}/sendMessage?parse_mode=markdown&chat_id=${chatId}&text=${text}`;

async function sendMessage() {
    try {
        while (true) {
            const response = await axios.get(url);
            if (response.status == 200) {
                console.log("Pesan terkirim:", text);
            }
        }
    } catch (error) {
        console.error("Gagal mengirim pesan:", error.message);
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

/*function spam() {
    sendMessage().then(spam);
}

spam();  */
