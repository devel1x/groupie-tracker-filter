function toggleConcertDates(element) {
    const cdates = element.nextElementSibling;
    cdates.classList.toggle('active');
    const icon = element.querySelector('.toggle-icon');
    icon.textContent = cdates.classList.contains('active') ? '-' : '+';
}
