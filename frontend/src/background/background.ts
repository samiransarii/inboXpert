interface EmailHeader {
  name: string;
  value: string;
}

interface GmailMessage {
  id: string;
  payload: {
    headers: EmailHeader[];
    parts?: {
      mimeType: string;
      body: {
        data?: string;
      };
    }[];
    body?: {
      data?: string;
    };
  };
}

interface Recipients {
  to: string[];
  cc: string[];
  bcc: string[];
}

interface Email {
  id: string;
  subject: string;
  body: string;
  sender: string;
  recipients: Recipients;
  headers: EmailHeader[];
}

interface SimplifiedEmail {
  id: string;
  subject: string;
  body: string;
  sender: string;
  recipients: string[];
  headers: Record<string, string>;
}

function decodeBase64(data: string): string {
  // Replace URL-safe characters and add padding
  const input = data.replace(/-/g, "+").replace(/_/g, "/");
  const pad = input.length % 4;
  const paddedInput = pad ? input + "=".repeat(4 - pad) : input;

  // Decode the string
  try {
    return decodeURIComponent(
      atob(paddedInput)
        .split("")
        .map(function (c) {
          return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
        })
        .join("")
    );
  } catch (e) {
    return atob(paddedInput); // Fallback to simple decode if UTF-8 fails
  }
}

function extractEmailBody(payload: GmailMessage["payload"]): string {
  // Check for body in the main payload
  if (payload.body?.data) {
    return decodeBase64(payload.body.data);
  }

  // Check for body in parts
  if (payload.parts) {
    // First try to find HTML content
    const htmlPart = payload.parts.find(
      (part) => part.mimeType === "text/html"
    );
    if (htmlPart?.body?.data) {
      return decodeBase64(htmlPart.body.data);
    }

    // Fall back to plain text
    const textPart = payload.parts.find(
      (part) => part.mimeType === "text/plain"
    );
    if (textPart?.body?.data) {
      return decodeBase64(textPart.body.data);
    }
  }

  return "";
}

function extractRecipients(headers: EmailHeader[]): Recipients {
  return {
    to:
      headers
        .find((h) => h.name === "To")
        ?.value?.split(",")
        .map((email) => email.trim())
        .filter((email) => email !== "") || [],
    cc:
      headers
        .find((h) => h.name === "Cc")
        ?.value?.split(",")
        .map((email) => email.trim())
        .filter((email) => email !== "") || [],
    bcc:
      headers
        .find((h) => h.name === "Bcc")
        ?.value?.split(",")
        .map((email) => email.trim())
        .filter((email) => email !== "") || [],
  };
}

async function getGmailInbox(): Promise<Email[]> {
  try {
    // Get OAuth token
    const { token } = await chrome.identity.getAuthToken({ interactive: true });

    // Fetch list of messages
    const response = await fetch(
      "https://www.googleapis.com/gmail/v1/users/me/messages?maxResults=50",
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    const data = await response.json();
    if (!data.messages) {
      throw new Error("No messages found in response");
    }

    // Fetch full message details for each email
    const emails = await Promise.all(
      data.messages.map(async (message: { id: string }) => {
        const messageDetails = await fetch(
          `https://www.googleapis.com/gmail/v1/users/me/messages/${message.id}?format=full`,
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );

        const messageData: GmailMessage = await messageDetails.json();
        if (!messageData.payload) {
          return {
            id: messageData.id,
            subject: "",
            body: "",
            sender: "",
            recipients: {
              to: [],
              cc: [],
              bcc: [],
            },
            headers: [],
          };
        }
        const headers = messageData.payload.headers;

        return {
          id: messageData.id,
          subject: headers.find((h) => h.name === "Subject")?.value || "",
          body: extractEmailBody(messageData.payload),
          sender: headers.find((h) => h.name === "From")?.value || "",
          recipients: extractRecipients(headers),
          headers: headers,
        };
      })
    );

    return emails;
  } catch (error) {
    console.error("Error fetching emails: ", error);
    throw error;
  }
}

// async function getGmailInbox(): Promise<Email[]> {
//   try {
//     // Get OAuth token
//     const { token } = await chrome.identity.getAuthToken({ interactive: true });

//     // Fetch list of messages
//     const response = await fetch(
//       "https://www.googleapis.com/gmail/v1/users/me/messages?maxResults=50",
//       {
//         headers: {
//           Authorization: `Bearer ${token}`,
//         },
//       }
//     );

//     const data = await response.json();
//     if (!data.messages) {
//       throw new Error("No messages found in response");
//     }

//     // Fetch full message details for each email
//     const emails = await Promise.all(
//       data.messages.map(async (message: { id: string }) => {
//         const messageDetails = await fetch(
//           `https://www.googleapis.com/gmail/v1/users/me/messages/${message.id}?format=full`,
//           {
//             headers: {
//               Authorization: `Bearer ${token}`,
//             },
//           }
//         );

//         const messageData: GmailMessage = await messageDetails.json();
//         const headers = messageData.payload.headers;

//         return {
//           id: messageData.id,
//           subject: headers.find((h) => h.name === "Subject")?.value || "",
//           body: extractEmailBody(messageData.payload),
//           sender: headers.find((h) => h.name === "From")?.value || "",
//           recipients: extractRecipients(headers),
//           headers: headers,
//         };
//       })
//     );

//     return emails;
//   } catch (error) {
//     console.error("Error fetching emails: ", error);
//     throw error;
//   }
// }

async function sendEmailsForCategorization(emails: Email[]): Promise<any> {
  // Convert Email[] to SimplifiedEmail[]
  const simplifiedEmails: SimplifiedEmail[] = emails.map(
    (email: Email): SimplifiedEmail => ({
      id: email.id,
      subject: email.subject,
      body: email.body,
      sender: email.sender,
      recipients: [
        ...email.recipients.to,
        ...email.recipients.cc,
        ...email.recipients.bcc,
      ],
      headers: email.headers.reduce(
        (
          acc: Record<string, string>,
          header: EmailHeader
        ): Record<string, string> => {
          acc[header.name] = header.value;
          return acc;
        },
        {}
      ),
    })
  );

  try {
    const response = await fetch("http://localhost:8080/categorize", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        emails: simplifiedEmails,
      }),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    console.log(response);
    console.log(response.json);
    return await response.json();
  } catch (error) {
    console.error("Error sending emails for categorization:", error);
    throw error;
  }
}

chrome.runtime.onMessage.addListener((request, _, sendResponse) => {
  if (request.action === "categorizeEmails") {
    getGmailInbox()
      .then(async (emails) => {
        try {
          const categorization = await sendEmailsForCategorization(emails);
          console.log(categorization);
          sendResponse({ success: true, categorization });
        } catch (error: unknown) {
          if (error instanceof Error) {
            sendResponse({ success: false, error: error.message });
          } else {
            sendResponse({
              success: false,
              error: "An unknown error occurred",
            });
          }
        }
      })
      .catch((error: unknown) => {
        if (error instanceof Error) {
          sendResponse({ success: false, error: error.message });
        } else {
          sendResponse({ success: false, error: "An unknown error occurred" });
        }
      });
    return true;
  }
});
