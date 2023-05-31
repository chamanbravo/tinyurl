const input = document.querySelector("input");
const button = document.querySelector("button");
const a = document.querySelector("a");

button.addEventListener("click", () => {
  fetch("/shorten", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      Url: input.value,
    }),
  })
    .then((res) => res.json())
    .then((data) => {
      a.innerText = `http://${window.location.host}/${data.tinyUrl}`;
      a.href = `http://${window.location.host}/${data.tinyUrl}`;
    });
});
