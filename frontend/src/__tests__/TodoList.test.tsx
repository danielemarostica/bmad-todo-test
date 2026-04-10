import { render, screen } from '@testing-library/react';
import { describe, it, expect } from 'vitest';
import TodoList from '../components/TodoList';

const mockTodos = [
  { id: '1', text: 'Buy groceries', completed: false, createdAt: '2026-04-10T12:00:00Z' },
  { id: '2', text: 'Call dentist', completed: true, createdAt: '2026-04-10T13:00:00Z' },
];

describe('TodoList', () => {
  it('renders todos when they exist', () => {
    render(
      <TodoList
        todos={mockTodos}
        loading={false}
        error={null}
        onToggle={() => {}}
        onDelete={() => {}}
      />
    );

    expect(screen.getByText('Buy groceries')).toBeInTheDocument();
    expect(screen.getByText('Call dentist')).toBeInTheDocument();
  });

  it('shows empty state when no todos exist', () => {
    render(
      <TodoList todos={[]} loading={false} error={null} onToggle={() => {}} onDelete={() => {}} />
    );

    expect(screen.getByText('No todos yet. Add one above.')).toBeInTheDocument();
  });

  it('shows loading state while fetching', () => {
    render(
      <TodoList todos={[]} loading={true} error={null} onToggle={() => {}} onDelete={() => {}} />
    );

    expect(screen.getByRole('status')).toHaveTextContent('Loading...');
  });

  it('shows error state when fetch fails', () => {
    render(
      <TodoList
        todos={[]}
        loading={false}
        error="Failed to load todos. Please try again."
        onToggle={() => {}}
        onDelete={() => {}}
      />
    );

    expect(screen.getByRole('alert')).toHaveTextContent('Failed to load todos');
  });

  it('visually distinguishes completed todos', () => {
    render(
      <TodoList
        todos={mockTodos}
        loading={false}
        error={null}
        onToggle={() => {}}
        onDelete={() => {}}
      />
    );

    const items = screen.getAllByRole('listitem');
    expect(items[0]).not.toHaveClass('completed');
    expect(items[1]).toHaveClass('completed');
  });
});
