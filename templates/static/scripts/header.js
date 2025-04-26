const searchInput = document.getElementById("blog-search-input");
const resultsContainer = document.getElementById("blog-search-results");
const resultsList = document.getElementById("results-list");
const emptyMessage = document.querySelector(".empty-message");
const closeSearchButton = document.getElementById("close-search-button");
const searchButton = document.querySelector(".search-button");
const header = document.querySelector(".header");
let blogs = []; // To store the fetched blog data

// Function to fetch the blog data from blogs.json
async function fetchBlogs() {
  try {
    const response = await fetch("blogs.json");
    if (!response.ok) {
      throw new Error(`Failed to fetch blogs.json: ${response.status}`);
    }
    blogs = await response.json();
    console.log("Fetched blogs:", blogs);
  } catch (error) {
    console.error("Error fetching blogs:", error);
    displayErrorMessage("Failed to load blog data. Please check back later.");
  }
}

// Function to display an error message
function displayErrorMessage(message) {
  const errorElement = document.createElement("div");
  errorElement.className = "error-message";
  errorElement.textContent = message;
  resultsContainer.parentElement.appendChild(errorElement);
}

// Function to handle the search input
function handleSearch() {
  const searchTerm = searchInput.value.toLowerCase().trim();

  // Clear previous results
  resultsList.innerHTML = "";

  if (searchTerm === "") {
    resultsContainer.classList.add("hidden");
    return;
  }

  // Filter blogs by search term
  const filteredBlogs = blogs.filter(
    (blog) =>
      blog.FolderName.toLowerCase().includes(searchTerm) ||
      blog.FrontMatterContent?.Title?.toLowerCase().includes(searchTerm),
  );

  if (filteredBlogs.length === 0) {
    resultsContainer.classList.remove("hidden");
    emptyMessage.classList.remove("hidden");
  } else {
    resultsContainer.classList.remove("hidden");
    emptyMessage.classList.add("hidden");

    // Create result items
    filteredBlogs.forEach((blog) => {
      const listItem = document.createElement("li");
      const link = document.createElement("a");
      link.href = `${blog.FolderName}/index.html`;
      link.textContent = blog.FrontMatterContent?.Title || blog.FolderName;
      listItem.appendChild(link);
      resultsList.appendChild(listItem);
    });
  }
}

// Function to handle closing the search results
function closeSearch() {
  resultsContainer.classList.add("hidden");
  searchInput.value = "";
  header.classList.remove("search-active");
}

// Event listener for the search input
searchInput.addEventListener("input", handleSearch);

// Event listener for the close button
closeSearchButton.addEventListener("click", closeSearch);

// Event listener for search button click (mobile)
searchButton.addEventListener("click", function () {
  header.classList.add("search-active");
  setTimeout(() => {
    searchInput.focus();
  }, 10);
});

// Event listener for clicks outside the search input and results
document.addEventListener("click", (event) => {
  if (
    !resultsContainer.contains(event.target) &&
    event.target !== searchInput &&
    event.target !== closeSearchButton &&
    event.target !== searchButton &&
    !searchButton.contains(event.target)
  ) {
    resultsContainer.classList.add("hidden");

    // Only remove active class if on mobile
    if (window.innerWidth <= 640) {
      header.classList.remove("search-active");
    }
  }
});

// Event listener for escape key
document.addEventListener("keydown", (event) => {
  if (event.key === "Escape") {
    closeSearch();
  }
});

// Initial setup: Fetch blogs
(async () => {
  await fetchBlogs();
})();
