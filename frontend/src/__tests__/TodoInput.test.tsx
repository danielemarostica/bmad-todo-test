import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, it, expect, vi } from 'vitest';
import TodoInput from '../components/TodoInput';

describe('TodoInput', () => {
  it('calls onSubmit with trimmed text when form is submitted', async () => {
    const onSubmit = vi.fn();
    render(<TodoInput onSubmit={onSubmit} />);

    const input = screen.getByPlaceholderText('What needs to be done?');
    await userEvent.type(input, 'Buy groceries');
    await userEvent.keyboard('{Enter}');

    expect(onSubmit).toHaveBeenCalledWith('Buy groceries');
  });

  it('clears the input after successful submit', async () => {
    render(<TodoInput onSubmit={() => {}} />);

    const input = screen.getByPlaceholderText('What needs to be done?');
    await userEvent.type(input, 'Buy groceries');
    await userEvent.keyboard('{Enter}');

    expect(input).toHaveValue('');
  });

  it('does not submit empty text', async () => {
    const onSubmit = vi.fn();
    render(<TodoInput onSubmit={onSubmit} />);

    await userEvent.keyboard('{Enter}');

    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('does not submit whitespace-only text', async () => {
    const onSubmit = vi.fn();
    render(<TodoInput onSubmit={onSubmit} />);

    const input = screen.getByPlaceholderText('What needs to be done?');
    await userEvent.type(input, '   ');
    await userEvent.keyboard('{Enter}');

    expect(onSubmit).not.toHaveBeenCalled();
  });
});
