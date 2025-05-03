package main

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow(
		"SELECT id,name,price FROM products WHERE id=$1",
		id,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func createProduct(product *Product) error {
	// INSERT INTO public.products(id, name, price) VALUES (?, ?, ?);

	_, err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)

	return err
}

func deleteProduct(id int) error {
	// DELETE FROM public.products WHERE <condition>;

	_, err := db.Exec(
		"DELETE FROM public.products WHERE id = $1;",
		id,
	)

	return err
}
