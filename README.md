 <h1>📚 Документация API заказов</h1>

<h2>💡 Примеры приложений</h2>

<h3>1. 📦 Панель управления заказами (Admin Dashboard)</h3>
  <p>Фронтенд-приложение, в котором администратор может:</p>
  <ul>
    <li>Просматривать список заказов</li>
    <li>Создавать, редактировать и удалять заказы</li>
    <li>Фильтровать заказы по статусу (pending, shipped, delivered)</li>
  </ul>

<h3>2. 🛍 Мини интернет-магазин</h3>
  <p>Клиентская часть, где пользователь может:</p>
  <ul>
    <li>Выбрать товар и оформить заказ</li>
    <li>Указать своё имя, количество и статус заказа</li>
    <li>Получать подтверждение заказа и отслеживать статус</li>
  </ul>

<h2>📌 Эндпоинты API</h2>
  <ul>
    <li><strong>GET</strong> <code>/orders/</code> — Получить список всех заказов</li>
    <li><strong>GET</strong> <code>/orders/:id</code> — Получить заказ по ID</li>
    <li><strong>POST</strong> <code>/orders/</code> — Создать новый заказ</li>
    <li><strong>PUT</strong> <code>/orders/:id</code> — Обновить заказ</li>
    <li><strong>DELETE</strong> <code>/orders/:id</code> — Удалить заказ</li>
  </ul>

<h2> Примеры запросов</h2>

<h3>Получить все заказы</h3>
  <pre><code>GET /orders/</code></pre>

<h3>Получить заказ по ID</h3>
  <pre><code>GET /orders/1</code></pre>

<h3>Создать заказ</h3>
  <pre><code>POST /orders/
Content-Type: application/json

{
  "customerName": "John Doe",
  "product": "iPhone 15",
  "quantity": 2,
  "status": "pending"
}
</code></pre>

<h3>Обновить заказ</h3>
  <pre><code>PUT /orders/1
Content-Type: application/json

{
  "customerName": "Jane Doe",
  "product": "MacBook Pro",
  "quantity": 1,
  "status": "shipped"
}
</code></pre>

<h3>Удалить заказ</h3>
  <pre><code>DELETE /orders/1</code></pre>

