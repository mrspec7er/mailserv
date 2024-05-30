import { simpleParser as parser, ParsedMail } from "mailparser";
import { SMTPServer, SMTPServerSession } from "smtp-server";

export interface ExtendedSMTPServerSession extends SMTPServerSession {
  mailFrom?: string;
  rcptTo?: string[];
}

export const Server = new SMTPServer({
  onData(
    stream,
    session: ExtendedSMTPServerSession,
    callback: (err?: Error | null) => void
  ) {
    parser(stream, {}, (err, parsedData) => {
      if (err) console.log("Error:", err);

      storeMessage(parsedData, session);
      // callback(new Error("Closing connection after processing email."));
      stream.on("end", callback);
    });
  },

  onMailFrom(address, session: ExtendedSMTPServerSession, callback) {
    session.mailFrom = address.address;
    callback();
  },

  onRcptTo(address, session: ExtendedSMTPServerSession, callback) {
    if (!session.rcptTo) {
      session.rcptTo = [];
    }
    session.rcptTo.push(address.address);
    callback();
  },
  disabledCommands: ["AUTH"],
});

export function storeMessage(
  data: ParsedMail,
  session: ExtendedSMTPServerSession
) {
  const payload = {
    subject: data.subject,
    sender: session.mailFrom,
    recipient: session.rcptTo.toString(),
    body: data.text,
  };
  fetch(`${process.env.REPOSITORY_DOMAIN}/emails`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  })
    .then((response) => response.json()) // Parse the JSON response
    .then((responseData) => {
      // console.log(responseData);
    })
    .catch((error) => {
      console.error(error);
    });
}
