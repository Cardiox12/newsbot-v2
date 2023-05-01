import IQueue from "@common/types/IQueue";
import Content from "@common/types/Content";
import amqp from "amqplib";


export class RabbitMQ implements IQueue {
    private addr: string;
    private name: string;
    private conn: amqp.Connection | null;
    private chan: amqp.Channel | null;

    constructor(name: string, addr: string) {
        this.name = name;
        this.addr = addr;
        this.conn = null;
        this.chan = null;
    }

    async connect(): Promise<void | Error> {
        this.conn = await amqp.connect(this.addr);
        this.chan = await this.conn.createChannel();
    }

    async publish(content: Content): Promise<void> {
        if ( !this.conn ){
            throw new Error("no connection has been made");
        }
        const contentBuffer = Buffer.from(JSON.stringify(content));

        await this.chan?.assertQueue(this.name);
        this.chan?.sendToQueue(this.name,  contentBuffer);
    }

    async get(): Promise<false | Content> {
        if ( !this.conn ){
            throw new Error("no connection has been made");
        }
        await this.chan?.assertQueue(this.name);
        
        const message = await this.chan?.get(this.name, { noAck: true });

        if (message){
            return JSON.parse(message.content.toString()) as Content;
        }
        return false;
    }

    async close(): Promise<void> {
        await this.chan?.close();
        await this.conn?.close();
    }
};