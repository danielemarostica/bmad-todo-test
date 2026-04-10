import type { Todo } from '../types/todo';
import TodoItem from './TodoItem';
import './TodoList.css';

interface TodoListProps {
  todos: Todo[];
  loading: boolean;
  error: string | null;
  onToggle: (id: string, completed: boolean) => void;
  onDelete: (id: string) => void;
}

export default function TodoList({ todos, loading, error, onToggle, onDelete }: TodoListProps) {
  if (loading) {
    return <div role="status" aria-label="Loading todos">Loading...</div>;
  }

  if (error) {
    return <div role="alert">{error}</div>;
  }

  if (todos.length === 0) {
    return <p>No todos yet. Add one above.</p>;
  }

  return (
    <ul className="todo-list" aria-label="Todo list">
      {todos.map((todo) => (
        <TodoItem key={todo.id} todo={todo} onToggle={onToggle} onDelete={onDelete} />
      ))}
    </ul>
  );
}
