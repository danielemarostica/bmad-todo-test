import { useState, useEffect, useCallback } from 'react';
import type { Todo } from './types/todo';
import * as api from './api/todos';
import TodoInput from './components/TodoInput';
import TodoList from './components/TodoList';
import './App.css';

export default function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchTodos = useCallback(async () => {
    try {
      setError(null);
      const data = await api.getTodos();
      setTodos(data);
    } catch {
      setError('Failed to load todos. Please try again.');
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchTodos();
  }, [fetchTodos]);

  const handleCreate = async (text: string) => {
    try {
      setError(null);
      const todo = await api.createTodo(text);
      setTodos((prev) => [...prev, todo]);
    } catch {
      setError('Failed to save. Please try again.');
    }
  };

  const handleToggle = async (id: string, completed: boolean) => {
    try {
      setError(null);
      const updated = await api.updateTodo(id, completed);
      setTodos((prev) => prev.map((t) => (t.id === id ? updated : t)));
    } catch {
      setError('Failed to update. Please try again.');
    }
  };

  const handleDelete = async (id: string) => {
    try {
      setError(null);
      await api.deleteTodo(id);
      setTodos((prev) => prev.filter((t) => t.id !== id));
    } catch {
      setError('Failed to delete. Please try again.');
    }
  };

  return (
    <main className="app">
      <h1>Todo App</h1>
      <TodoInput onSubmit={handleCreate} />
      <TodoList
        todos={todos}
        loading={loading}
        error={error}
        onToggle={handleToggle}
        onDelete={handleDelete}
      />
    </main>
  );
}
