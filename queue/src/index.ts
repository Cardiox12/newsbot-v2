import { RabbitMQ } from "./rabbit/rabbit";
import Content from "@common/types/Content";


async function main() {
    const rabbit: RabbitMQ = new RabbitMQ("newsbot", "amqp://localhost");

    await rabbit.connect();

    for (let i = 0; i < 5; ++i) {
        await rabbit.publish({
            url: `http://news/${i}`,
            title: `title ${i}`,
            source: "news"
        });
    }

    setTimeout(async () => {
        while (true) {
            const message: Content | false = await rabbit.get();

            if (message) {
                console.table(message);
            } else {
                break;
            }
        }
        await rabbit.close();
    }, 5000);
}

main();
