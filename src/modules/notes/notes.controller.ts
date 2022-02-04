import { Controller, Get } from '@nestjs/common'
import { NotesService } from './notes.service'

@Controller('notes')
export class NotesController {
  constructor(private readonly notesService: NotesService) {}

  @Get()
  getHello(): string {
    return 'Hello World'
  }
}
