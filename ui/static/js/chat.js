const chatBox = document.getElementById("chatBox");
const form = document.getElementById("form");
const msg = document.getElementById("msg");

form.addEventListener("submit", async (e) => {
	e.preventDefault();

	if (!msg.value.trim()) {
		return;
	}

	try {
		const response = await fetch(
			`/messages/${room}`,
			{
				method: "POST",
				headers: {
					"Content-Type": "application/x-www-form-urlencoded",
				},
				body: `message=${encodeURIComponent(msg.value)}`
			}
		);

		if (!response.ok) {
			console.error("POST failed:", response.status);
			return;
		}

		msg.value = "";

		loadMessages();

	} catch (err) {
		console.error(err);
	}
});

async function loadMessages() {
	try {
		const res = await fetch(`/messages/${room}`);
		const text = await res.text();

		chatBox.innerHTML = "";

		text.trim().split("\n").forEach(line => {
			if (!line) return;

			const parts = line.split("|");

			if (parts.length !== 2) return;

			const username = parts[0];
			const message = parts.slice(1).join("|"); 

			const div = document.createElement("div");
			div.classList.add("message");

			div.innerHTML = `
				<span class="user">${username}</span>
				<span class="content">${message}</span>
			`;

			chatBox.appendChild(div);
		});
		chatBox.scrollTop = chatBox.scrollHeight;

	} catch (err) {
		console.error("Error loading messages:", err);
	}
}

loadMessages();

setInterval(loadMessages, 5000);