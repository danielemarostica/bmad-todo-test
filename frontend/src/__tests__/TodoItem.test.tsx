import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, it, expect, vi } from 'vitest';
import TodoItem from '../components/TodoItem';

const activeTodo = {
  id: '1',
  text: 'Buy groceries',
  completed: false,
  createdAt: '2026-04-10T12:00:00Z',
};

const completedTodo = {
  id: '2',
  text: 'Call dentist',
  completed: true,
  createdAt: '2026-04-10T13:00:00Z',
};

describe('TodoItem', () => {
  it('renders todo text', () => {
    render(
      <TodoItem todo={activeTodo} onToggle={() => {}} onDelete={() => {}} />
    );

    expect(screen.getByText('Buy groceries')).toBeInTheDocument();
  });

  it('calls onToggle with id and new completed state when checkbox clicked', async () => {
    const onToggle = vi.fn();
    render(
      <TodoItem todo={activeTodo} onToggle={onToggle} onDelete={() => {}} />
    );

    const checkbox = screen.getByRole('checkbox');
    await userEvent.click(checkbox);

    expect(onToggle).toHaveBeenCalledWith('1', true);
  });

  it('calls onToggle to uncomplete when completed todo checkbox clicked', async () => {
    const onToggle = vi.fn();
    render(
      <TodoItem todo={completedTodo} onToggle={onToggle} onDelete={() => {}} />
    );

    const checkbox = screen.getByRole('checkbox');
    await userEvent.click(checkbox);

    expect(onToggle).toHaveBeenCalledWith('2', false);
  });

  it('calls onDelete with id when delete button clicked', async () => {
    const onDelete = vi.fn();
    render(
      <TodoItem todo={activeTodo} onToggle={() => {}} onDelete={onDelete} />
    );

    const deleteBtn = screen.getByRole('button', { name: /delete/i });
    await userEvent.click(deleteBtn);

    expect(onDelete).toHaveBeenCalledWith('1');
  });

  it('shows checkbox as checked for completed todo', () => {
    render(
      <TodoItem todo={completedTodo} onToggle={() => {}} onDelete={() => {}} />
    );

    expect(screen.getByRole('checkbox')).toBeChecked();
  });

  it('shows checkbox as unchecked for active todo', () => {
    render(
      <TodoItem todo={activeTodo} onToggle={() => {}} onDelete={() => {}} />
    );

    expect(screen.getByRole('checkbox')).not.toBeChecked();
  });
});
