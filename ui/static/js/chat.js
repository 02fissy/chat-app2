
let user = localStorage.getItem("userID");

if (!user) {
    user = Math.floor(Math.random() * 1000000).toString();
    localStorage.setItem("userID", user);
}

const chatBox = document.getElementById("chatBox");
const form = document.getElementById("form");
const msg = document.getElementById("msg");


form.addEventListener("submit", async (e) => {
	e.preventDefault();

	if (!msg.value.trim()) return;

	const data = new FormData();
	data.append("message", msg.value);

	await fetch(`/messages/${room}/users/${user}`, {
		method: "POST",
		body: data
	});

	msg.value = "";
	loadMessages();
});

async function loadMessages() {

	const res = await fetch(`/messages/${room}/data`);
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
}
loadMessages();
setInterval(loadMessages, 5000);