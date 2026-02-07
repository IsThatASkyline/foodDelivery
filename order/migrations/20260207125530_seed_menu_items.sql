-- +goose Up
-- +goose StatementBegin
INSERT INTO orders.menu_items (id, name, price) VALUES
    (1, 'Маргарита',        45000),
    (2, 'Пепперони',        52000),
    (3, 'Четыре сыра',      56000),
    (4, 'Гавайская',        53000),
    (5, 'Цезарь с курицей', 48000),
    (6, 'Бургер классик',   49000),
    (7, 'Картофель фри',    19000),
    (8, 'Наггетсы',         26000),
    (9, 'Кола 0.5',         15000),
    (10,'Чизкейк',          32000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
