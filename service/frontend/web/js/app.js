const BASE_URL = 'http://localhost:8080';

document.getElementById('recommendationForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const customerId = document.getElementById('customerId').value;
    if (customerId) {
        fetch(`${BASE_URL}/recommendations?customerId=${encodeURIComponent(customerId)}`)
            .then(response => response.json())
            .then(data => {
                if (data.length > 0) {
                    let table = `<table class='table table-bordered mt-3'>
                        <thead>
                            <tr>
                                <th>Title</th>
                                <th>Description</th>
                                <th>Product</th>
                            </tr>
                        </thead>
                        <tbody>`;
                    data.forEach(offer => {
                        table += `<tr>
                            <td>${offer.title}</td>
                            <td>${offer.description}</td>
                            <td>${offer.product}</td>
                        </tr>`;
                    });
                    table += `</tbody></table>`;
                    document.getElementById('response').innerHTML = table;
                } else {
                    document.getElementById('response').innerHTML = `<p class='text-warning'>No offers available.</p>`;
                }
            })
            .catch(error => {
                document.getElementById('response').innerHTML = `<p class='text-danger'>Error fetching recommendations.</p>`;
                console.error('Error:', error);
            });
    }
});