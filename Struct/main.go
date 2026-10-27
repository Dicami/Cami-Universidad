package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func main() {
	inventario := Inventario{
		Productos: make(map[int]Producto),
	}

	product1 := NuevoProducto(1, "Celular", 300, 4)
	product2 := NuevoProducto(2, "Comedor", 10000, 20)
	product3 := NuevoProducto(3, "Television", 150, 10)

	inventario.AgregarProducto(product1)
	inventario.AgregarProducto(product2)
	inventario.AgregarProducto(product3)

	fmt.Println("Lista de prodcutos")
	inventario.ListarProductos()

	fmt.Println("Buscar products")
	producto, existe := inventario.BuscarProducto(2)

	if existe {
		producto.MostrarInfo()
	} else {
		fmt.Println("Producto no encontrado")
	}

	fmt.Println("Actualizar")
	product1.ActualizarPrecio(700.00)
	product1.ActualizarCantidad(8)
	inventario.AgregarProducto(product1)

	product1.MostrarInfo()

	fmt.Println("Eliminar producto")
	inventario.EliminarProducto(3)

	fmt.Println("Inventario final con el prodcuto eliminado")
	inventario.ListarProductos()

}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{id, nombre, precio, cantidad}
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	p.Cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Println("ID:", p.ID)
	fmt.Println("Nombre:", p.Nombre)
	fmt.Println("Precio:", p.Precio)
	fmt.Println("Cantidad:", p.Cantidad)

}

type Inventario struct {
	Productos map[int]Producto
}

func (i *Inventario) AgregarProducto(p Producto) {
	i.Productos[p.ID] = p
}

func (i *Inventario) EliminarProducto(id int) {
	delete(i.Productos, id)
}

func (i Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := i.Productos[id]
	return p, existe
}

func (i Inventario) ListarProductos() {
	for _, p := range i.Productos {
		p.MostrarInfo()
	}
}
