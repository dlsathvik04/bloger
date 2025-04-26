const searchInput = document.getElementById("blog-search-input");
const resultsContainer = document.getElementById("blog-search-results");
const resultsList = document.getElementById("results-list");
const emptyMessage = document.getElementById("empty-message");

let blogs = [];

// Fetch blog data
async function fetchBlogs() {
  try {
    const response = await fetch("/blogs.json");
    if (!response.ok) throw new Error("Failed to fetch blogs.json");
    blogs = await response.json();
  } catch (error) {
    console.error(error);
  }
}

// Handle search input
function handleSearch() {
  const searchTerm = searchInput.value.trim().toLowerCase();
  resultsList.innerHTML = "";

  if (!searchTerm) {
    resultsContainer.classList.remove("show");
    return;
  }

  const filteredBlogs = blogs.filter(
    (blog) =>
      blog.FolderName.toLowerCase().includes(searchTerm) ||
      blog.FrontMatterContent?.Title?.toLowerCase().includes(searchTerm),
  );

  if (filteredBlogs.length === 0) {
    resultsContainer.classList.add("show");
    emptyMessage.classList.remove("d-none");
  } else {
    emptyMessage.classList.add("d-none");
    filteredBlogs.forEach((blog) => {
      const item = document.createElement("li");
      item.innerHTML = `<a href="/${blog.FolderName}/index.html" class="dropdown-item">${blog.FrontMatterContent?.Title || blog.FolderName}</a>`;
      resultsList.appendChild(item);
    });
    resultsContainer.classList.add("show");
  }
}

// Close dropdown on outside click
document.addEventListener("click", (e) => {
  if (!searchInput.contains(e.target) && !resultsContainer.contains(e.target)) {
    resultsContainer.classList.remove("show");
  }
});

// Hide results on Escape key
document.addEventListener("keydown", (e) => {
  if (e.key === "Escape") {
    resultsContainer.classList.remove("show");
    searchInput.value = "";
  }
});

// Event listeners
searchInput.addEventListener("input", handleSearch);

// Initial fetch
fetchBlogs();
