import Content from "./Content";

export default interface IQueue {
    publish(content: Content) : void;
    get() : Promise<false | Content>;
    connect() : Promise<void | Error>;
    close() : Promise<void>;
}