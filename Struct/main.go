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

	var id int
	var nombre string
	var precio float64
	var cantidad int

	fmt.Println("Ingresar prodcutos")

	for i := 0; i < 3; i++ {
		fmt.Println("Producto", i+1)

		fmt.Print("Ingrese el ID: ")
		fmt.Scan(&id)

		fmt.Print("Ingrese el nombre: ")
		fmt.Scan(&nombre)

		fmt.Print("Ingrese el precio: ")
		fmt.Scan(&precio)

		fmt.Print("Ingrese la cantidad: ")
		fmt.Scan(&cantidad)

		producto := NuevoProducto(id, nombre, precio, cantidad)
		inventario.AgregarProducto(producto)
	}

	fmt.Println("Lista de prodcutos")
	inventario.ListarProductos()

	fmt.Println("Buscar productos")
	fmt.Print("Ingrese el ID del producto que desea buscar: ")
	fmt.Scan(&id)

	producto, existe := inventario.BuscarProducto(id)

	if existe {
		producto.MostrarInfo()
	} else {
		fmt.Println("Producto no encontrado")
	}

	fmt.Println("Actualizar productos")
	fmt.Print("Ingrese el ID del producto que desea actualizar: ")
	fmt.Scan(&id)

	producto, existe = inventario.BuscarProducto(id)

	if existe {
		fmt.Print("Ingrese el nuevo precio: ")
		fmt.Scan(&precio)

		fmt.Print("Ingrese la nueva cantidad: ")
		fmt.Scan(&cantidad)

		producto.ActualizarPrecio(precio)
		producto.ActualizarCantidad(cantidad)

		inventario.AgregarProducto(producto)

		fmt.Println("Producto actualizado:")
		producto.MostrarInfo()
	} else {
		fmt.Println("Producto no encontrado")
	}

	fmt.Println("Eliminar prodcucto")
	fmt.Print("Ingrese el ID del producto que desea eliminar: ")
	fmt.Scan(&id)

	inventario.EliminarProducto(id)

	fmt.Println("Inventario con el producto eliminado")
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
