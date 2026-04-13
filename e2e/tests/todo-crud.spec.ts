import { test, expect } from '@playwright/test';

// Clean up todos before each test via API
test.beforeEach(async ({ request }) => {
  const res = await request.get('http://localhost:8080/api/v1/todos');
  const todos = await res.json();
  for (const todo of todos) {
    await request.delete(`http://localhost:8080/api/v1/todos/${todo.id}`);
  }
});

test.describe('Todo CRUD', () => {
  test('shows empty state on first load', async ({ page }) => {
    await page.goto('/');
    await expect(page.getByText('No todos yet. Add one above.')).toBeVisible();
  });

  test('creates a new todo', async ({ page }) => {
    await page.goto('/');
    await page.getByPlaceholder('What needs to be done?').fill('Buy groceries');
    await page.getByPlaceholder('What needs to be done?').press('Enter');

    await expect(page.getByText('Buy groceries')).toBeVisible();
    // Input should be cleared
    await expect(page.getByPlaceholder('What needs to be done?')).toHaveValue('');
  });

  test('completes and uncompletes a todo', async ({ page, request }) => {
    // Create a todo via API
    await request.post('http://localhost:8080/api/v1/todos', {
      data: { text: 'Test toggle' },
    });

    await page.goto('/');
    const checkbox = page.getByRole('checkbox', { name: /Test toggle/ });

    // Complete
    await checkbox.click();
    await expect(checkbox).toBeChecked();

    // Uncomplete
    await checkbox.click();
    await expect(checkbox).not.toBeChecked();
  });

  test('deletes a todo', async ({ page, request }) => {
    await request.post('http://localhost:8080/api/v1/todos', {
      data: { text: 'To be deleted' },
    });

    await page.goto('/');
    await expect(page.getByText('To be deleted')).toBeVisible();

    await page.getByRole('button', { name: /Delete "To be deleted"/ }).click();

    await expect(page.getByText('To be deleted')).not.toBeVisible();
    await expect(page.getByText('No todos yet. Add one above.')).toBeVisible();
  });

  test('persists todos across page reload', async ({ page, request }) => {
    await request.post('http://localhost:8080/api/v1/todos', {
      data: { text: 'Persistent todo' },
    });

    await page.goto('/');
    await expect(page.getByText('Persistent todo')).toBeVisible();

    // Reload
    await page.reload();
    await expect(page.getByText('Persistent todo')).toBeVisible();
  });
});
