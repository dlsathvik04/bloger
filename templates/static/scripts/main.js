async function loadHTML(selector, url) {
  const res = await fetch(url);
  const html = await res.text();
  document.querySelector(selector).innerHTML = html;
}

async function loadBlogs() {
  const res = await fetch("blogs.json");
  const blogs = await res.json();
  const cardTemplate = await fetch("blog_card.html").then((r) => r.text());

  const container = document.getElementById("blog-container");
  container.innerHTML = "";

  blogs.forEach((blog) => {
    const card = cardTemplate
      .replace("{{title}}", blog.name)
      .replace("{{description}}", `This is a blog about ${blog.name}`)
      .replace("{{link}}", `${blog.name}/index.html`);
    const wrapper = document.createElement("div");
    wrapper.innerHTML = card;
    container.appendChild(wrapper);
  });
}

window.addEventListener("DOMContentLoaded", async () => {
  await loadHTML("#header", "header.html");
  await loadHTML("#footer", "footer.html");
  await loadBlogs();
});
