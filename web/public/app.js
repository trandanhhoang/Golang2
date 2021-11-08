window.addEventListener("DOMContentLoaded", (_) => {
  let websocket = new WebSocket("ws://" + window.location.host + "/websocket");
  let room = document.getElementById("chat-text");

  websocket.addEventListener("message", function (e) {
    let data = JSON.parse(e.data);
    // let chatContent = `<p><strong>${data.username}</strong>: ${data.text}</p>`;
    var chatContent = document.createElement("p");

    var user = document.createElement("strong");
    user.textContent = `${data.username}`;
    chatContent.append(user);
    chatContent.append(` : ${data.text}`);
    room.appendChild(chatContent);
    // room.append(addLi);
    room.scrollTop = room.scrollHeight; // Auto scroll to the bottom
  });

  let form = document.getElementById("input-form");
  form.addEventListener("submit", function (event) {
    event.preventDefault();
    let username = document.getElementById("input-username");
    let text = document.getElementById("input-text");
    websocket.send(
      JSON.stringify({
        username: username.value,
        text: text.value,
      })
    );
    text.value = "";
  });
});
