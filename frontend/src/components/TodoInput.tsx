import { useState } from 'react';
import './TodoInput.css';

interface TodoInputProps {
  onSubmit: (text: string) => void;
}

export default function TodoInput({ onSubmit }: TodoInputProps) {
  const [text, setText] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const trimmed = text.trim();
    if (!trimmed) return;
    onSubmit(trimmed);
    setText('');
  };

  return (
    <form onSubmit={handleSubmit} className="todo-input">
      <label htmlFor="todo-text" className="sr-only">Add a todo</label>
      <input
        id="todo-text"
        type="text"
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="What needs to be done?"
        aria-label="New todo text"
      />
      <button type="submit">Add</button>
    </form>
  );
}
