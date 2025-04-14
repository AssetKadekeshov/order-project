<h2>📌 Модель заказа (Order)</h2>
<pre>
{
  "id": 1,
  "customerName": "Иван Иванов",
  "products": ["Product A", "Product B"],
  "totalQuantity": 2,
  "status": "created"
}
</pre>

<h2>📡 API-запросы</h2>
<p>Все запросы отправляются на <code>/orders</code>.</p>

<h3>🔍 Получить все заказы</h3>
<pre>GET /orders/</pre>
<p><strong>Параметры фильтрации (опционально):</strong></p>
<ul>
  <li><code>status</code>: фильтр по статусу (created, shipped и т.п.)</li>
  <li><code>product</code>: фильтр по товару</li>
  <li><code>search</code>: поиск по имени клиента</li>
</ul>

<p><strong>📌 Пример:</strong></p>
<pre>GET /orders/?status=created&product=Product A&search=Иван</pre>

<h3>📄 Получить заказ по ID</h3>
<pre>GET /orders/{id}</pre>
<p><strong>📌 Пример:</strong></p>
<pre>GET /orders/1</pre>

<h3>➕ Создать заказ</h3>
<pre>
POST /orders/
Content-Type: application/json

{
  "customerName": "Новый клиент",
  "products": ["Product X", "Product Y"],
  "status": "pending"
}
</pre>

<h3>✏️ Обновить заказ</h3>
<pre>
PUT /orders/{id}
Content-Type: application/json

{
  "customerName": "Обновлённый клиент",
  "products": ["Product X"],
  "status": "shipped"
}
</pre>

<h3>❌ Удалить заказ</h3>
<pre>DELETE /orders/{id}</pre>

<h3>🔄 Автоматическая миграция</h3>
<p>В <code>main.go</code> при запуске вызывается:</p>
<pre>db.AutoMigrate(&models.Order{})</pre>
<p>Это автоматически создаёт таблицу <code>orders</code>, если её нет в базе данных.</p>

<h3>🧠 Бизнес-логика</h3>
<p>Бизнес-логика отделена в слое <code>service</code>. Все действия (создание, фильтрация, обновление) проходят через <code>OrderService</code>, который взаимодействует с репозиторием (<code>OrderRepositoryImpl</code>), реализующим работу с базой.</p>

<h3>🛠 Примеры фильтрации</h3>
<p><strong>📌 Вызов:</strong></p>
<pre>GET /orders?status=created&product=Milk&search=Алихан</pre>
