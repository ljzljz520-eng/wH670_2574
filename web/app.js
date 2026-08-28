const status = document.getElementById('status');
function show(message){ status.textContent = message; }
function render(records){ show(`${records.length} records available`); }
window.examArchive = { show, render };
