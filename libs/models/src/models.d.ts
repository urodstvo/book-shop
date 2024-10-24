declare module "models" {
  type Book = {
    id: number;
    name: string;
    cover: string;
    author: string;
    rating: number | null;
    rating_count: number;
    annotation: string;
    price: number;
    pageCount: number;
    orders_count: number;
    stock_count: number;
    published_by: string;
    published_at: Date;
    created_at: Date;
    updated_at: Date;
  };

  type Order = {
    id: number;
    user_id: number;
    payment_id: number;
    price: number;
    status: OrderStatus;
    created_at: Date;
    updated_at: Date;
  };

  type Genre = {
    id: number;
    name: string;
  };

  type Cart = {
    id: number;
    user_id: number;
    book_id: number;
    quantity: number;
    created_at: Date;
    updated_at: Date;
  };

  type User = {
    id: number;
    login: string;
    password: string;
    name: string;
    role: string;
    created_at: Date;
    updated_at: Date;
  };

  type Payment = {
    id: number;
    user_id: number;
    card_number: string;
    card_type: string;
    cardholder_name: string;
    card_expired_at: Date;
    created_at: Date;
    updated_at: Date;
  };

  export { Book, Order, Genre, Cart, User, Payment };
}
