package service

// /new XXX YYY ZZZ` - создаёт новую задачу
func (s Service) NewTask() {
	// s.SendMessage(ctx, &bot.SendMessageParams{
	// 	ChatID: update.Message.Chat.ID,
	// 	Text:   fmt.Sprintf(env.comment),
	// })
	// return

	// parts := strings.SplitN(update.Message.Text, " ", 4) // Разбиваем сообщение на части
	// if len(parts) < 4 {
	// 	b.SendMessage(ctx, &bot.SendMessageParams{
	// 		ChatID: update.Message.Chat.ID,
	// 		Text:   "Используйте: /new [ID] [Title] [Description]",
	// 	})
	// 	return
	// }

	// taskID := parts[1]
	// taskTitle := parts[2]
	// taskDesc := parts[3]

	// fmt.Println("newTaskHandler:", taskID, taskTitle, taskDesc)

	// tasks[taskID] = &Task{
	// 	ID:    taskID,
	// 	Title: taskTitle,
	// }

	// b.SendMessage(ctx, &bot.SendMessageParams{
	// 	ChatID: update.Message.Chat.ID,
	// 	Text:   fmt.Sprintf("Задача создана: ID=%s, Title=%s", taskID, taskTitle),
	// })
}
