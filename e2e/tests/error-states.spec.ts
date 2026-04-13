import { test, expect } from '@playwright/test';

test.beforeEach(async ({ request }) => {
  const res = await request.get('http://localhost:8080/api/v1/todos');
  const todos = await res.json();
  for (const todo of todos) {
    await request.delete(`http://localhost:8080/api/v1/todos/${todo.id}`);
  }
});

test.describe('Input validation', () => {
  test('does not create a todo with empty text', async ({ page }) => {
    await page.goto('/');
    await page.getByPlaceholder('What needs to be done?').press('Enter');

    // Should still show empty state
    await expect(page.getByText('No todos yet. Add one above.')).toBeVisible();
  });

  test('does not create a todo with whitespace-only text', async ({ page }) => {
    await page.goto('/');
    await page.getByPlaceholder('What needs to be done?').fill('   ');
    await page.getByPlaceholder('What needs to be done?').press('Enter');

    await expect(page.getByText('No todos yet. Add one above.')).toBeVisible();
  });
});
