import amqp from 'amqplib';
import dotenv from 'dotenv';

export default class Queue {
  rabbitMQUrl: string;

  constructor() {
    dotenv.config();
    this.rabbitMQUrl = process.env.RABBITMQ_URL || '';
  }

  async publish(queue: string, data: any) {
    const connection = await amqp.connect(this.rabbitMQUrl);
    const channel = await connection.createChannel();
    channel.assertQueue(queue, { durable: true });
    channel.sendToQueue(queue, Buffer.from(JSON.stringify(data)));
  }

  async consume(queue: string, callback: Function) {
    const connection = await amqp.connect(this.rabbitMQUrl);
    const channel = await connection.createChannel();
    channel.assertQueue(queue, { durable: true });
    channel.consume(queue, async (msg: any) => {
      const input = JSON.parse(msg.content.toString());
      try {
        await callback(input);
        channel.ack(msg);
      } catch (e: any) {
        console.log(e);
      }
    });
  }
}
