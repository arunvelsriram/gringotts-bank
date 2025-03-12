const BASE_URL = 'http://localhost:8080';

document.getElementById('recommendationForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const customerId = document.getElementById('customerId').value;
    if (customerId) {
        fetch(`${BASE_URL}/recommendations?customerId=${encodeURIComponent(customerId)}`)
            .then(response => response.json())
            .then(data => {
                let recommendations = data.recommendations;
                console.log(`recommendations count: ${recommendations.length}`)
                if (recommendations.length > 0) {
                    let table = `
                    <h5> Customer </h5>
                    <p>
                        <b>ID:</b> ${data.customerId} / <b>Name:</b> ${data.customerName} / <b>Age:</b> ${data.customerAge}
                    </p>
                    <h5> Recommendations </h5>
                    <table class='table table-bordered mt-3'>
                        <thead>
                            <tr>
                                <th>Title</th>
                                <th>Description</th>
                                <th>Product</th>
                            </tr>
                        </thead>
                        <tbody>`;
                    recommendations.forEach(recommendation => {
                        table += `<tr>
                            <td>${recommendation.title}</td>
                            <td>${recommendation.description}</td>
                            <td>${recommendation.product}</td>
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