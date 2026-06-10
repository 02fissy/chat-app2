let user = sessionStorage.getItem("userID");

if (!user) {
	user = Math.floor(Math.random() * 100000);
	sessionStorage.setItem("userID", user);
}

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
			`/messages/${room}/users/${user}`,
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

			const userID = parts[0];
			const message = parts[1];

			const div = document.createElement("div");

			div.textContent = `User ${userID}: ${message}`;

			chatBox.appendChild(div);
		});

	} catch (err) {
		console.error(err);
	}
}

loadMessages();

setInterval(loadMessages, 5000);